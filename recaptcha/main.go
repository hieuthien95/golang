package main

import (
	"context"
	"fmt"

	recaptcha "cloud.google.com/go/recaptchaenterprise/apiv1"
	"google.golang.org/api/option"
	recaptchapb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
)

func main() {
	// TODO(developer): Replace these variables before running the sample.
	projectID := "caphe"
	recaptchaSiteKey := "<site_key>"
	token := "<response_token>"
	recaptchaAction := "register"

	createAssessment(projectID, recaptchaSiteKey, token, recaptchaAction)
}

/**
 * Create an assessment to analyze the risk of an UI action.
 *
 * @param projectID: GCloud Project ID
 * @param recaptchaSiteKey: Site key obtained by registering a domain/app to use recaptcha services.
 * @param token: The token obtained from the client on passing the recaptchaSiteKey.
 * @param recaptchaAction: Action name corresponding to the token.
 */
func createAssessment(projectID string, recaptchaSiteKey string, token string, recaptchaAction string) {

	// Create the recaptcha client.
	ctx := context.Background()

	saKey := "conf/sa-caphe.json"
	client, err := recaptcha.NewClient(ctx, option.WithCredentialsFile(saKey))
	if err != nil {
		fmt.Printf("Error creating reCAPTCHA client\n")
	}
	defer client.Close()

	// Set the properties of the event to be tracked.
	event := &recaptchapb.Event{
		Token:   token,
		SiteKey: recaptchaSiteKey,
	}

	assessment := &recaptchapb.Assessment{
		Event: event,
	}

	// Build the assessment request.
	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: assessment,
		Parent:     fmt.Sprintf("projects/%s", projectID),
	}

	response, err := client.CreateAssessment(
		ctx,
		request)

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	// Check if the token is valid.
	if response.TokenProperties.Valid == false {
		fmt.Printf("The CreateAssessment() call failed because the token"+
			" was invalid for the following reasons: %v",
			response.TokenProperties.InvalidReason)
		return
	}

	// Check if the expected action was executed.
	if response.TokenProperties.Action == recaptchaAction {
		// Get the risk score and the reason(s).
		// For more information on interpreting the assessment,
		// see: https://cloud.google.com/recaptcha-enterprise/docs/interpret-assessment
		fmt.Printf("The reCAPTCHA score for this token is:  %v",
			response.RiskAnalysis.Score)

		for _, reason := range response.RiskAnalysis.Reasons {
			fmt.Printf(reason.String() + "\n")
		}
		return
	}

	fmt.Printf("The action attribute in your reCAPTCHA tag does " +
		"not match the action you are expecting to score")
}
