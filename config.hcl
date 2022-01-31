source = ["./alfred-workflow_kaba-ts.alfredworkflow"]
bundle_id = "com.kaba-tech.alfred-ts"

apple_id {
  username = "yourAppleID"
  password = "yourApp-SpecificPassword"
  provider = "yourProviderShortname"
}

sign {
  application_identity = "yourDeveloperIDApplicationHash like 8D3983A55C3D535941D89814FC0DAAC958D2A50F"
}

zip {
  output_path = "alfred-workflow_kaba-ts.alfredworkflow.notarize.zip"
}