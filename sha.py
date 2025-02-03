from github import Github, InputGitTreeElement
import os

# Authenticate
g = Github(os.environ['GITHUB_TOKEN'])
repo = g.get_repo("databricks/terraform-provider-databricks")

# Create blobs for 3 files
blobs = [
    repo.create_git_blob(content=f"123", encoding="utf-8"),
    repo.create_git_blob(content=f"123", encoding="utf-8"),
    repo.create_git_blob(content=f"123", encoding="utf-8")
]

# Create tree elements
tree_elements = [
    InputGitTreeElement(path="file1.txt", mode="100644", type="blob", sha=blobs[0].sha),
    InputGitTreeElement(path="subdir/file2.txt", mode="100644", type="blob", sha=blobs[1].sha),
    InputGitTreeElement(path="file3.txt", mode="100644", type="blob", sha=blobs[2].sha)
]

# Create tree and commit
head_ref = repo.get_git_ref("heads/hectorcast-db/test-commit-2")
base_tree = repo.get_git_tree(sha=head_ref.object.sha)
new_tree = repo.create_git_tree(tree_elements, base_tree)
parent_commit = repo.get_git_commit(head_ref.object.sha)

new_commit = repo.create_git_commit(
    message="Add multiple files",
    tree=new_tree,
    parents=[parent_commit]
)

# Update branch reference
head_ref.edit(new_commit.sha)