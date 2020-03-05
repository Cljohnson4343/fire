package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch resources from a Firestore instance",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := firestore.NewClient(ctx, "")
		if err != nil {
			fmt.Sprintf("unable to initialize Firestore client: %c\n", err)
		}

		fmt.Println("All issues:")
		iter := client.Collection("issuesDevelopment").Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("unable to retreive document: %v\n", err)

			}
			fmt.Println(doc.Data())
		}

	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
