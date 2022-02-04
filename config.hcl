# https://github.com/mitchellh/gon

source = ["./alfred-workflow_kaba-ts.alfredworkflow"]
bundle_id = "com.kaba-tech.alfred-ts"

apple_id {
  username = "" # Your AppleID
  password = "@keychain:AC_PASSWORD" # Your AppleSpecificPassword https://support.apple.com/en-al/HT204397
  provider = "" # Run `xcrun altool --list-providers -u "AppleID" -p "AppleSpecificPassword"` to find your provider shortname
}

sign {
  application_identity = "yourDeveloperIDApplicationHash like 8D3983A55C3D535941D89814FC0DAAC958D2A50F" # Run `security find-identity -v` to find your type:[Developer ID Application] identity
}

zip {
  output_path = "alfred-workflow_kaba-ts.alfredworkflow.notarize.zip"
}