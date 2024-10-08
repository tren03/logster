package azureblob

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

var NUM = 0
var client *azblob.Client
var flag = false
var containerName = "test"

// to remove all data once docker container stops
// docker run --rm -p 10000:10000 --name azurite-container mcr.microsoft.com/azure-storage/azurite azurite-blob --blobHost 0.0.0.0

func CreateContainer() {

	// only run for the first time
	if flag == false {
		connectionString := "DefaultEndpointsProtocol=http;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=http://azurite:10000/devstoreaccount1;"
		// Create a client to connect to Azurite
		azclient, err := azblob.NewClientFromConnectionString(connectionString, nil)
		if err != nil {
			log.Println("err connecting")
		}
		client = azclient
		flag = true

	}

	fmt.Println("creating container")
	containerCreateResp, err := client.CreateContainer(context.TODO(), containerName, nil)
	if err != nil {
		log.Println("err making testcontainer", err)
	}
	fmt.Println(containerCreateResp)
}

func DownloadAllBlob() {
	file, err := os.OpenFile("./logs.txt", os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Println("error opening file", err)
	}
	for i := 0; i < NUM; i++ {
		blobName := fmt.Sprintf("LOGFILE%d", i)
		decodedData := DownloadBlob(blobName)
		_, err = file.Write([]byte(decodedData))
		if err != nil {
			log.Println("error writing to file", err)
		}

	}
}

func DownloadBlob(blobName string) string {
	blobDownloadResponse, err := client.DownloadStream(context.TODO(), containerName, blobName, nil)
	if err != nil {
		log.Println("err in downloading blob")
	}
	reader := blobDownloadResponse.Body
	downloadData, err := io.ReadAll(reader)
	if err != nil {
		log.Println("err in reading downloaded data")
	}
	err = reader.Close()
	if err != nil {
		log.Println("reader close issue")
	}
	fmt.Println("done with download")

	return string(downloadData)

}
func UploadToBlob(logArray string) {
	fmt.Println("WE ARE IN THE UPLOAD DATA FUNC")
	// ===== 2. Upload =====
	blobData := logArray
	blobName := fmt.Sprintf("LOGFILE%d", NUM)
	NUM += 1

	fmt.Println("uploading blob")
	uploadResp, err := client.UploadStream(context.TODO(),
		containerName,
		blobName,
		strings.NewReader(blobData),
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"Foo": to.Ptr("Bar")},
			Tags:     map[string]string{"Year": "2022"},
		})

	if err != nil {
		log.Println("err upload")
	}
	fmt.Println(uploadResp)
	fmt.Println("WE FINISHED UPLOADING THE DATA")
}

func GetBlobInfo() {

	var totalsize int64
	fmt.Println("WE ARE GETTING INFO ABOUT BLOB")
	pager := client.NewListBlobsFlatPager(containerName, nil)

	// Continue fetching pages until no more remain
	for pager.More() {
		// Advance to the next page
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Printf("Error while enumerating blobs: %v\n", err)
			return
		}

		// Iterate over blobs in the current page
		for _, blob := range page.Segment.BlobItems {
			// Get blob name
			blobName := *blob.Name

			// Get blob size (in bytes)
			blobSize := int64(0)
			if blob.Properties.ContentLength != nil {
				blobSize = *blob.Properties.ContentLength
				totalsize += blobSize
			}

			// Print blob name and size
			fmt.Printf("Blob Name: %s, Size: %d bytes\n", blobName, blobSize)
		}
		fmt.Println("TOTAL BLOB SIZE : ", totalsize)
	}
	fmt.Println("WE FINISHED")
}
