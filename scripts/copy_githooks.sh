#!/bin/sh

cp ./githooks/prepare-commit-msg.sh .git/hooks/prepare-commit-msg
cp ./githooks/pre-commit.sh .git/hooks/pre-commit
cp ./githooks/pre-push.sh .git/hooks/pre-push
chmod 755 .git/hooks/prepare-commit-msg
chmod 755 .git/hooks/pre-commit
chmod 755 .git/hooks/pre-push