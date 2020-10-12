package consumer

import "github.com/Azure/azure-storage-blob-go/azblob"

func main() {
	println("consumer")
	println(azblob.PremiumPageBlobAccessTierNone)
}
