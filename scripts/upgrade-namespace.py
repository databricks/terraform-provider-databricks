import glob, re, sys

old_provider = "databrickslabs/databricks"
new_provider = "databricks/databricks"

files_to_fix = []
regex = re.compile(r"source\s+=\s+\"{}\"".format(old_provider))

for terraform_configuration_file in glob.glob('**/**.tf', recursive=True):
    contents = open(terraform_configuration_file, 'r').read()
    if not regex.findall(contents):
        continue
    print(f'[+] File {terraform_configuration_file} matches.')
    files_to_fix.append((terraform_configuration_file, contents))

if not len(files_to_fix):
    print("NOTHING: Didn't find any mentions of the databricks provider.")
    sys.exit(1)

if 'yes' == input(f'Type yes to confirm fix of {len(files_to_fix)} files: ').lower():
    for terraform_configuration_file, contents in files_to_fix:
        with open(terraform_configuration_file, 'w') as f:
            new_contents = regex.sub(f'source = "{new_provider}"', contents)
            f.write(new_contents)
    print(f'SUCCESS: Fixed {len(files_to_fix)} files!')
else: print("ABORT: Didn't receive 'yes'.")
