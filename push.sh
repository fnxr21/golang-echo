git add .

echo 'Enter the commit message:'

read commitMessage

git commit -m "$commitMessage"

git push main dev

echo 'Thank you for updated your repository'