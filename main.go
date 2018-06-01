package main

import (
	"log"
	"os"
	"fmt"
	// "path/filepath"
	"flag"
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Name              string     `json:"name"`
	Id 								int        `json:"id"`
	Description       string     `json:"description"`
	SimpleDescription string     `json:"simple_description"`
	TileId            int        `json:"tile_id"`
	DirectoryIndex    int        `json:"directory_index"`
}

/*
 simple script to create json objects for tiles from directory/image information
*/
func main() {
	/*
   * flags for creating json object from cli
   * - readDir <read_dir> = directory where script should find files one would like to jsonify
   * - jsonFileName <json_file_name> = output file name one would like jsonified objects to be written to
	 */

	// flag start
	readDir := flag.String("read_dir", "/Users/azim/Documents/2DGAMEASSETS/DungeonCrawl32px/player/base", "directory where script should find files one would like to jsonify")
	jsonFileName := flag.String("json_file_name", "output.json", "output file name one would like jsonified stuff to be written to")

	flag.Parse()
	// flag end

	
	files, err := ioutil.ReadDir(*readDir)
  if err != nil {
    log.Fatal(err)
  }
  var jsonDocumentString = ""
  log.Printf("directory count: %d", len(files))
  for index, f := range files {
    fmt.Printf("[%d]: %s", index, f.Name())
    fmt.Println()
    // tmpFileName := fmt.Sprintf("%s/%s",searchDir,f.Name())
    var tmpItem = Item{}
		tmpItem.Name = f.Name()
		tmpItem.Id = 0
		tmpItem.Description = ""
		tmpItem.SimpleDescription = ""
		tmpItem.TileId = 0
		tmpItem.DirectoryIndex = index


    itemJson, _ := json.MarshalIndent(tmpItem,"", "  ")
    switch index {
    	case 0:
    		jsonDocumentString += fmt.Sprintf("[%s\n,", itemJson)
    		break
    	case (len(files) - 1):
    		jsonDocumentString += fmt.Sprintf("%s]\n", itemJson)
    		break
    	default:
    		jsonDocumentString += fmt.Sprintf("%s,\n", itemJson)
    		break
    }
    

  }
	
  writingFile, err := os.OpenFile(*jsonFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
	    panic(err)
	}

	defer writingFile.Close()

	if _, err = writingFile.WriteString(jsonDocumentString); err != nil {
	    panic(err)
	}
}
