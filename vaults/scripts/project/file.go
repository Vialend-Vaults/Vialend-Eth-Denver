package include

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

const fn = "C:\\Users\\xdotk\\torukmakto\\vialend\\contracts\\vaults\\v2.0\\scripts\\file\\ini"

//const cfg_file = "C:\\Users\\xdotk\\sendbox\\leonopteryx\\vialend\\contracts\\vaults\\v2.1\\scripts\\file\\config.json"
//const cfg_file = "C:\\Users\\xdotk\\torukmakto\\vialend\\contracts\\vaults\\v2.1\\scripts\\file\\config.json"

//const cfg_file = "../scripts/file/config.json"
const cfg_file = "C:\\Users\\xdotk\\sandbox\\leonopteryx\\vialend\\contracts\\vaults\\v2.1\\scripts\\file\\config.json"

type Configuration struct {
	Description string `json:"Description"`
	Contracts   struct {
		STRAT_DEPLOYER string
		VAULT_DEPLOYER string
		VAULT_FACTORY  string
		VAULT_STRATEGY string
		VAULT          string
		VAULT_BRIDGE   string
		CALLEE         string
		SWAPHELPER     string
	} `json:"Contracts"`
}

var Cfg Configuration

func ConfigWrite() {

	var Cfgs2 Configuration

	var result map[string]interface{}
	Load(&result)

	sid := strconv.Itoa(Networkid)
	mapstructure.Decode(result[sid].(map[string]interface{}), &Cfgs2)
	if !reflect.DeepEqual(Cfg, Cfgs2) {
		result[sid] = &Cfg
		Save(cfg_file, &result)
	}

}
func ConfigParser() {
	var result map[string]interface{}
	Load(&result)

	sid := strconv.Itoa(Networkid)

	fmt.Println(result[sid])
	if result[sid] == nil {
		return
	}
	mapstructure.Decode(result[sid].(map[string]interface{}), &Cfg)

}

func getFile(relfile string) string {

	f, _ := filepath.Abs(relfile)

	ff := strings.Replace(f, "\\", "\\\\", -1)

	return ff
}

// Load gets your config from the json file,
// and fills your struct with the option
func Load(o interface{}) {

	filename := getFile(cfg_file)
	if b, err := ioutil.ReadFile(filename); err == nil {
		err = json.Unmarshal(b, &o)
	} else {
		log.Fatal("ReadFile:", err)
		fmt.Println("here:", filename)

	}

}

// Save will save your struct to the given filename,
// this is a good way to create a json template
func Save(filename string, o interface{}) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(filename, j, 0660)
	}
	return err
}

func ReadINI() []string {

	settings := ReadFile(fn)

	if len(settings) < 3 {
		log.Fatal("check ini file. invalid settings")
	}

	return (settings)
}

func WriteINI(msgs []string) {

	WriteFile(fn, msgs, false)
}

//fn "../../file/tick"
func ReadFile(fi string) []string {

	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("ReadFile:Open: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	sln, e := Readln(r)

	fileInArray := []string{}

	for e == nil {

		s := strings.TrimSpace(sln)

		if len(s) > 0 {
			fileInArray = append(fileInArray, s) // a == [3 4]
		}
		sln, e = Readln(r)
	}

	return fileInArray

}

// fn "../file/tick"
func WriteFile(fn string, msgs []string, isAppend bool) {

	var buf []byte

	for i := 0; i < len(msgs); i++ {
		ss := []byte(msgs[i] + "\n")
		buf = append(buf, ss...)
	}

	if isAppend {
		d0, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal("WriteFileAppend:Readfile err:", err)
		}

		s := strings.TrimSpace(string(d0))

		buf = append(buf, s...)
	}
	err := os.WriteFile(fn, buf, 0644)
	if err != nil {
		log.Fatal("writefile err:", err)
	}

}
