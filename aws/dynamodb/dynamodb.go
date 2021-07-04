package dynamodb

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createServiceClient() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("ap-northeast-1"),
			Endpoint: aws.String("http://localhost:8000"),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new session: %w", err)
	}

	svc := dynamodb.New(sess)

	return svc, nil
}

func prepareTable(svc *dynamodb.DynamoDB, tableName string) error {
	ok, err := existsTable(svc, tableName)
	if err != nil {
		return fmt.Errorf("failed to describe table: %w", err)
	}

	if ok {
		fmt.Printf("%s table already active\n", tableName)
		return nil
	}

	err = createTable(svc, tableName)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	input := &dynamodb.DescribeTableInput{TableName: aws.String(tableName)}

	err = svc.WaitUntilTableExists(input)
	if err != nil {
		return fmt.Errorf("failed to wait until table exists: %w", err)
	}

	fmt.Printf("%s table has been created\n", tableName)

	return nil
}

func existsTable(svc *dynamodb.DynamoDB, tableName string) (bool, error) {
	input := &dynamodb.DescribeTableInput{TableName: aws.String(tableName)}

	output, err := svc.DescribeTable(input)
	if err != nil {
		var aerr awserr.Error
		if errors.As(err, &aerr) {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return false, nil
			default:
				return false, aerr
			}
		}
	}

	if *output.Table.TableStatus == "ACTIVE" {
		return true, nil
	}

	return false, nil
}

func createTable(svc *dynamodb.DynamoDB, tableName string) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("build_id"),
				AttributeType: aws.String("S"),
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("build_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		StreamSpecification: &dynamodb.StreamSpecification{
			StreamEnabled:  aws.Bool(true),
			StreamViewType: aws.String("NEW_AND_OLD_IMAGES"),
		},
		TableName: aws.String(tableName),
	}

	if _, err := svc.CreateTable(input); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func printTables(svc *dynamodb.DynamoDB) {
	output, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Tables:")

	if len(output.TableNames) == 0 {
		fmt.Println("  not found")
		return
	}

	for _, name := range output.TableNames {
		fmt.Println("  " + *name)
	}
}

func Execute() {
	const TableName = "builds"

	svc, err := createServiceClient()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	err = prepareTable(svc, TableName)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	printTables(svc)
}
