package commentator

import "testing"

const goFileTxt = "{\n    \"content\": \"package main\\n func main(){\\n println('hi')\\n}\"\n}\n"

func TestConvertInputDataToStruct(t *testing.T) {
	t.Run("Get Content", func(t *testing.T) {
		_, err := ConvertInputDataToStruct(goFileTxt)
		if err != nil {
			t.Fatal("Could not get the contents")
		}
	})
}

func TestParser_Exec(t *testing.T) {
	t.Run("Successful Parse Exec", func(t *testing.T) {

	})
}


