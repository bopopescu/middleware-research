package vod

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// TranscodeTemplate is a nested struct in vod response
type TranscodeTemplate struct {
	TranscodeTemplateId  string   `json:"TranscodeTemplateId" xml:"TranscodeTemplateId"`
	Video                string   `json:"Video" xml:"Video"`
	Audio                string   `json:"Audio" xml:"Audio"`
	Container            string   `json:"Container" xml:"Container"`
	MuxConfig            string   `json:"MuxConfig" xml:"MuxConfig"`
	TransConfig          string   `json:"TransConfig" xml:"TransConfig"`
	Definition           string   `json:"Definition" xml:"Definition"`
	EncryptSetting       string   `json:"EncryptSetting" xml:"EncryptSetting"`
	PackageSetting       string   `json:"PackageSetting" xml:"PackageSetting"`
	SubtitleList         string   `json:"SubtitleList" xml:"SubtitleList"`
	OpeningList          string   `json:"OpeningList" xml:"OpeningList"`
	TailSlateList        string   `json:"TailSlateList" xml:"TailSlateList"`
	TemplateName         string   `json:"TemplateName" xml:"TemplateName"`
	TranscodeFileRegular string   `json:"TranscodeFileRegular" xml:"TranscodeFileRegular"`
	Clip                 string   `json:"Clip" xml:"Clip"`
	Rotate               string   `json:"Rotate" xml:"Rotate"`
	WatermarkIds         []string `json:"WatermarkIds" xml:"WatermarkIds"`
}
