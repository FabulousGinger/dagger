package tasks

import "context"

func Pipeline(ctx context.Context) (err error) {

	err = Test(ctx)
	CheckIfError(err)

	err = Scan(ctx)
	CheckIfError(err)

	err = Terraform(ctx, "apply")
	CheckIfError(err)

	err = ECSDeploy(ctx)
	CheckIfError(err)

	return
}
