package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main(){
	//bundle is their translation god-object.  it keeps track of loaded translation files, default language &c
	bundle := i18n.NewBundle(language.English)

	//sets how goi18n should handle unmarshaling different file types. the default is toml, so we make sure to tell
	//it what to do with json
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	//look up the message file, parse it using our unmarshal function we registered above,
	//and store the translations in memory in the bundle
	bundle.LoadMessageFile("active.es.json")
	bundle.LoadMessageFile("active.en.json")

	//localizer is our workhorse for performing the real-time translation substitutions.
	//it is initialized with a reference to the bundle of loaded translations and a target language to look up
	//the lookup should match the naming of the files loaded into the bundle.  If we had saved our translations
	//in a file named active.es-SP.json, we would need to supply "es-SP" here rather than simply "es"
	localizer := i18n.NewLocalizer(bundle, "es")

	//templateData is a map of the arguments to insert.  their formatting strings are done in the text/template style, eg
	//"Hola, {{.Name}}"
	templateData := map[string]string{ "Name": "Steve" }

	//we call localize passing a pointer to a config object which defines the message to lookup and the map to do
	//substitutions from.  it feels a bit odd to not also be passing in the target language for translations, but the
	//localizer is already initialized with a target language
	//it will return an error if it can't find the message ID entry for that language
	espanol, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:        "HI",
		TemplateData:   templateData,
	})
	if err !=nil{
		fmt.Println("omg it broke.")
		fmt.Println(err)
	}else{
		fmt.Println(espanol)
	}
}
