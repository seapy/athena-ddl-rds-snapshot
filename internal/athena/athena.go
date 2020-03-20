package athena

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func Update(sql string, workgroup string) (*athena.StartQueryExecutionOutput, error) {
	session := session.Must(session.NewSession(&aws.Config{}))
	ath := athena.New(session)
	in := athena.StartQueryExecutionInput{
		QueryString: &sql,
		WorkGroup: &workgroup,
	}
	return ath.StartQueryExecution(&in)
}
