#!/usr/bin/env python3
"""Track and retest Terraform acceptance tests against the in-process fakebricks.

Source of truth: tests.jsonl (one record per acceptance test). Each record holds
the test's package, level, last result, and — for failures — the classified
cause and the precise endpoint/message to fix on the fakebricks side.

Workflow:
  1. Fix something in fakebricks (fakebricks-tf branch).
  2. Retest the affected set, e.g. all tests blocked by an endpoint:
        python3 fakebricks-acc/retest.py run --match "/scim/v2/Users"
     Anything that now passes is flipped to status=PASS / fix_status=fixed.
  3. Regenerate the human view:
        python3 fakebricks-acc/retest.py report

Selectors (AND-combined; run/list/report all accept them):
  --match SUBSTR   substring match over "package/test cause endpoint detail"
  --package PKG    exact package dir (e.g. scim, catalog)
  --cause CAUSE    cause bucket (e.g. 501_NOT_IMPLEMENTED, WORKSPACE_ID)
  --status S       current status (FAIL, SKIP, PASS, TIMEOUT)
  --open           only fix_status=open (not yet fixed) — default for `run`
  --test ID        exact "package/test"
Use `list`/`--dry-run` to preview a selection without running it.
"""
import argparse, json, os, re, subprocess, sys, datetime

HERE = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(HERE)
STORE = os.path.join(HERE, "tests.jsonl")
REPORT = os.path.join(HERE, "report.md")
TF = os.environ.get("TF_ACC_TERRAFORM_PATH", "/usr/local/bin/terraform")
TODAY = datetime.date.today().isoformat()

# ----- classification ---------------------------------------------------------

def decisive(log):
    out = []
    for ln in log.splitlines():
        if re.search(r'well-known|databricks-config', ln):
            continue
        m = re.search(r'diagnostic_summary="([^"]*)"', ln)
        if m: out.append(m.group(1)); continue
        m = re.search(r'Step \d+/\d+ error: (.*)', ln)
        if m: out.append(m.group(1)); continue
        if re.search(r'(^|\| )Error:', ln): out.append(re.sub(r'^\s*\|?\s*', '', ln))
        if re.search(r'Not equal|Error Trace|Should |Expected', ln): out.append(ln.strip())
    seen, uniq = set(), []
    for x in out:
        if x not in seen: seen.add(x); uniq.append(x)
    return " ".join(uniq)[:300]

def classify(log, code):
    """Return (status, cause, endpoint, detail)."""
    if "test timed out after" in log or code == 124:
        return "TIMEOUT", "TIMEOUT", "", decisive(log)[:160]
    if "no tests to run" in log:
        return "NORUN", "NORUN", "", ""
    failed = ("--- FAIL" in log) or (code not in (0,))
    if "--- SKIP" in log and not failed:
        lines = log.splitlines()
        msg = ""
        for i, ln in enumerate(lines):
            if ln.strip().startswith("--- SKIP") and i > 0:
                msg = re.sub(r'^[A-Za-z0-9_]+\.go:\d+:\s*', '', lines[i-1].strip())
                break
        return "SKIP", "SKIP", "", msg[:120]
    if not failed and ("--- PASS" in log or code == 0):
        return "PASS", "", "", ""
    d = decisive(log)
    def ep(pat):
        m = re.search(pat, log)
        return re.sub(r'/accounts/[^/]+/', '/accounts/{acct}/', m.group(0)) if m else ""
    if re.search(r'is not implemented|NOT_IMPLEMENTED', d):
        endpoint = ep(r'[A-Z]+ /api/[^\s"]+(?= is not implemented)')
        return "FAIL", "501_NOT_IMPLEMENTED", endpoint, d
    if re.search(r'no fakebricks handler|ENDPOINT_NOT_FOUND', d):
        return "FAIL", "404_NO_HANDLER", ep(r'no fakebricks handler for [A-Z]+ /[^\s"]+'), d
    if re.search(r'non-refresh plan was not empty|inconsistent result after apply|unexpected new value', d):
        return "FAIL", "POST_APPLY_DRIFT", "", d
    if re.search(r'invalid character|cannot unmarshal|json: ', d):
        return "FAIL", "REQUEST_PARSE", "", d
    if "DBC extraction is not modeled" in d:
        return "FAIL", "NOTEBOOK_DBC", "", d
    if "Group has no role" in d:
        return "FAIL", "GROUP_ROLES", "", d
    if re.search(r"doesn't exist|does not exist|DOES_NOT_EXIST", d):
        return "FAIL", "NOT_SEEDED", "", d
    if "workspace_id" in d:
        return "FAIL", "WORKSPACE_ID", "", d
    if re.search(r'Not equal|Error Trace|Should |Expected|assert', d):
        return "FAIL", "CHECK_ASSERTION", "", d
    return "FAIL", "OTHER", "", d

# ----- store ------------------------------------------------------------------

def load():
    if not os.path.exists(STORE): return []
    return [json.loads(l) for l in open(STORE) if l.strip()]

def save(recs):
    recs.sort(key=lambda r: (r["status"] != "FAIL", r.get("cause",""), r["id"]))
    with open(STORE, "w") as f:
        for r in recs: f.write(json.dumps(r) + "\n")

def env_for(level):
    e = dict(os.environ)
    # -mod=mod so the local `replace github.com/databricks/cli => .../fakebricks-tf`
    # is read live from the worktree. Without it, a populated vendor/ (created by
    # `make build`/`make lint`) would pin a stale snapshot and hide fakebricks fixes.
    e.update(CLOUD_ENV="aws", DATABRICKS_FAKE="1", TF_ACC_TERRAFORM_PATH=TF, GOFLAGS="-mod=mod")
    if level == "ACCT":
        e["TEST_ENVIRONMENT_TYPE"] = "ACCOUNT"; e["ACCT_ID"] = "fake-acct-0000"
        e["DATABRICKS_ACCOUNT_ID"] = "fake-acct-0000"
    else:
        e["TEST_ENVIRONMENT_TYPE"] = "UC_WORKSPACE"
    return e

def run_test(rec, timeout):
    pkg, func, level = rec["package"], rec["test"], rec.get("level","WS")
    try:
        p = subprocess.run(
            ["timeout", str(timeout), "go", "test", "-v", "-count=1",
             "-run", f"^{func}$", f"./{pkg}/", "-timeout", f"{timeout-10}s"],
            cwd=REPO, env=env_for(level),
            stdout=subprocess.PIPE, stderr=subprocess.STDOUT, universal_newlines=True)
        return classify(p.stdout, p.returncode)
    except Exception as ex:
        return "FAIL", "RUNNER_ERROR", "", str(ex)[:160]

# ----- selection --------------------------------------------------------------

def select(recs, a):
    out = []
    for r in recs:
        blob = f"{r['id']} {r.get('cause','')} {r.get('endpoint','')} {r.get('detail','')}"
        if a.match and a.match.lower() not in blob.lower(): continue
        if a.package and r["package"] != a.package: continue
        if a.cause and r.get("cause") != a.cause: continue
        if a.status and r["status"] != a.status: continue
        if a.test and r["id"] != a.test: continue
        if getattr(a, "open", False) and r.get("fix_status") != "open": continue
        out.append(r)
    return out

# ----- report -----------------------------------------------------------------

def report(recs):
    import collections
    fails = [r for r in recs if r["status"] in ("FAIL","TIMEOUT") and r.get("fix_status")!="fixed"]
    fixed = [r for r in recs if r.get("fix_status")=="fixed"]
    skips = [r for r in recs if r["status"]=="SKIP"]
    passes = [r for r in recs if r["status"]=="PASS" and r.get("fix_status")!="fixed"]
    L = []
    L.append("# Fakebricks acceptance status\n")
    L.append(f"_Updated {TODAY}._ Open failures: **{len(fails)}** · Fixed: **{len(fixed)}** · "
             f"Passing: **{len(passes)}** · Skipped (gated): **{len(skips)}**.\n")
    L.append("Retest a set after a fakebricks fix, e.g. `python3 fakebricks-acc/retest.py run --match \"/scim/v2/Users\"`.\n")
    by = collections.defaultdict(list)
    for r in fails: by[(r.get("cause",""), r.get("endpoint",""))].append(r)
    L.append("## Open failures by cause\n")
    for (cause, ep), items in sorted(by.items(), key=lambda kv: -len(kv[1])):
        head = f"{cause}" + (f" — `{ep}`" if ep else "")
        L.append(f"### {head} ({len(items)})")
        for r in sorted(items, key=lambda r: r["id"]):
            note = f" — {r['detail']}" if (not ep and r.get("detail")) else ""
            L.append(f"- [ ] {r['id']}{note}")
        L.append("")
    if fixed:
        L.append("## Fixed ✅\n")
        for r in sorted(fixed, key=lambda r: r["id"]):
            L.append(f"- [x] {r['id']} (was {r.get('cause','?')}; fixed {r.get('fixed_at','')})")
        L.append("")
    L.append("## Skipped — gated, not a fakebricks failure (grouped)\n")
    sg = collections.defaultdict(list)
    for r in skips: sg[r.get("detail","skip")].append(r["id"])
    for reason, ids in sorted(sg.items(), key=lambda kv: -len(kv[1])):
        L.append(f"- **{reason}** ({len(ids)})")
    open(REPORT, "w").write("\n".join(L) + "\n")
    print(f"wrote {REPORT}: {len(fails)} open, {len(fixed)} fixed, {len(skips)} skipped")

# ----- commands ---------------------------------------------------------------

def cmd_build(a):
    """Build tests.jsonl from a finished sweep: exact names from --from-tsv
    (pkg<TAB>func<TAB>result<TAB>reason), cause re-derived from --results-dir logs."""
    rd = a.results_dir
    recs = []
    for ln in open(a.from_tsv):
        parts = ln.rstrip("\n").split("\t")
        if len(parts) < 3: continue
        pkg, func = parts[0], parts[1]
        safe = f"{pkg}__{func}".replace("/", "_")
        logp = os.path.join(rd, safe + ".log")
        if os.path.exists(logp):
            log = open(logp, errors="ignore").read()
            code = 0 if (("--- PASS" in log or "--- SKIP" in log) and "--- FAIL" not in log) else 1
            status, cause, ep, detail = classify(log, code)
        else:
            status, cause, ep, detail = parts[2], "", "", (parts[3] if len(parts) > 3 else "")
        level = "ACCT" if func.startswith("TestMwsAcc") else ("UC" if func.startswith("TestUcAcc") else "WS")
        recs.append(dict(id=f"{pkg}/{func}", package=pkg, test=func, level=level,
                         status=status, cause=cause, endpoint=ep, detail=detail,
                         fix_status=("open" if status in ("FAIL","TIMEOUT") else status.lower()),
                         last_run=TODAY))
    save(recs)
    print(f"built {STORE}: {len(recs)} records")
    report(recs)

def cmd_list(a):
    recs = load(); sel = select(recs, a)
    for r in sel:
        print(f"{r['status']:<7} {r.get('fix_status',''):<7} {r['id']}  [{r.get('cause','')}] {r.get('endpoint','')}")
    print(f"\n{len(sel)} selected")

def cmd_run(a):
    recs = load()
    # With no selector, default to retesting everything still open (unfixed).
    if not (a.match or a.package or a.cause or a.status or a.test): a.open = True
    sel = select(recs, a)
    # Don't rerun already-passing tests unless the user explicitly asked for PASS.
    sel = [r for r in sel if r["status"] != "PASS" or a.status == "PASS"]
    if a.dry_run:
        for r in sel: print(f"would run: {r['id']}  [{r.get('cause','')}] {r.get('endpoint','')}")
        print(f"\n{len(sel)} would run"); return
    print(f"retesting {len(sel)} test(s), timeout={a.timeout}s each\n")
    newly_fixed, still = [], []
    for r in sel:
        status, cause, ep, detail = run_test(r, a.timeout)
        prev = r["status"]
        r["last_run"] = TODAY
        if status == "PASS":
            if prev != "PASS":
                r["status"] = "PASS"; r["fix_status"] = "fixed"; r["fixed_at"] = TODAY
                newly_fixed.append(r["id"]); mark = "✅ FIXED"
            else:
                r["status"] = "PASS"; mark = "pass"
        else:
            r["status"], r["cause"], r["endpoint"], r["detail"] = status, cause, ep, detail
            still.append(r["id"]); mark = f"{status} [{cause}]"
        print(f"  {mark:<22} {r['id']}")
    save(recs)
    print(f"\nnewly fixed: {len(newly_fixed)} | still failing: {len(still)}")
    report(recs)

def cmd_report(a):
    report(load())

def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = ap.add_subparsers(dest="cmd")
    sub.required = True
    def add_sel(p):
        p.add_argument("--match"); p.add_argument("--package"); p.add_argument("--cause")
        p.add_argument("--status"); p.add_argument("--test"); p.add_argument("--open", action="store_true")
    b = sub.add_parser("build"); b.add_argument("--results-dir", required=True)
    b.add_argument("--from-tsv", required=True); b.set_defaults(fn=cmd_build)
    l = sub.add_parser("list"); add_sel(l); l.set_defaults(fn=cmd_list)
    r = sub.add_parser("run"); add_sel(r); r.add_argument("--timeout", type=int, default=240)
    r.add_argument("--dry-run", action="store_true"); r.set_defaults(fn=cmd_run)
    rp = sub.add_parser("report"); rp.set_defaults(fn=cmd_report)
    a = ap.parse_args(); a.fn(a)

if __name__ == "__main__":
    main()
