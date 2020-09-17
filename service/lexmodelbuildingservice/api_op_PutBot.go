// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates an Amazon Lex conversational bot or replaces an existing bot. When you
// create or update a bot you are only required to specify a name, a locale, and
// whether the bot is directed toward children under age 13. You can use this to
// add intents later, or to remove intents from an existing bot. When you create a
// bot with the minimum information, the bot is created or updated but Amazon Lex
// returns the  response FAILED. You can build the bot after you add one or more
// intents. For more information about Amazon Lex bots, see how-it-works (). If you
// specify the name of an existing bot, the fields in the request replace the
// existing values in the $LATEST version of the bot. Amazon Lex removes any fields
// that you don't provide values for in the request, except for the
// idleTTLInSeconds and privacySettings fields, which are set to their default
// values. If you don't specify values for required fields, Amazon Lex throws an
// exception.  <p>This operation requires permissions for the
// <code>lex:PutBot</code> action. For more information, see
// <a>security-iam</a>.</p>
func (c *Client) PutBot(ctx context.Context, params *PutBotInput, optFns ...func(*Options)) (*PutBotOutput, error) {
	stack := middleware.NewStack("PutBot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutBotMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutBotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBot(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutBot",
			Err:           err,
		}
	}
	out := result.(*PutBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBotInput struct {
	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation. A user interaction session remains active for the amount of time
	// specified. If no conversation occurs during this time, the session expires and
	// Amazon Lex deletes any data provided before the timeout. For example, suppose
	// that a user chooses the OrderPizza intent, but gets sidetracked halfway through
	// placing an order. If the user doesn't complete the order within the specified
	// time, Amazon Lex discards the slot information that it gathered, and the user
	// must start over. If you don't include the idleSessionTTLInSeconds element in a
	// PutBot operation request, Amazon Lex uses the default value. This is also true
	// if the request replaces an existing bot. The default is 300 seconds (5 minutes).
	IdleSessionTTLInSeconds *int32
	// If you set the processBehavior element to BUILD, Amazon Lex builds the bot so
	// that it can be run. If you set the element to SAVE Amazon Lex saves the bot, but
	// doesn't build it. If you don't specify this value, the default value is BUILD.
	ProcessBehavior types.ProcessBehavior
	// When set to true a new numbered version of the bot is created. This is the same
	// as calling the CreateBotVersion operation. If you don't specify createVersion,
	// the default is false.
	CreateVersion *bool
	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to a
	// website, program, or other application that is directed or targeted, in whole or
	// in part, to children under age 13 and subject to COPPA. By specifying false in
	// the childDirected field, you confirm that your use of Amazon Lex is not related
	// to a website, program, or other application that is directed or targeted, in
	// whole or in part, to children under age 13 and subject to COPPA. You may not
	// specify a default value for the childDirected field that does not accurately
	// reflect whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. If your use of Amazon Lex relates to a
	// website, program, or other application that is directed in whole or in part, to
	// children under age 13, you must obtain any required verifiable parental consent
	// under COPPA. For information regarding the use of Amazon Lex in connection with
	// websites, programs, or other applications that are directed or targeted, in
	// whole or in part, to children under age 13, see the Amazon Lex FAQ.
	// (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool
	// When Amazon Lex can't understand the user's input in context, it tries to elicit
	// the information a few times. After that, Amazon Lex sends the message defined in
	// abortStatement to the user, and then aborts the conversation. To set the number
	// of retries, use the valueElicitationPrompt field for the slot type. For example,
	// in a pizza ordering bot, Amazon Lex might ask a user "What type of crust would
	// you like?" If the user's response is not one of the expected responses (for
	// example, "thin crust, "deep dish," etc.), Amazon Lex tries to elicit a correct
	// response a few more times. For example, in a pizza ordering application,
	// OrderPizza might be one of the intents. This intent might require the CrustType
	// slot. You specify the valueElicitationPrompt field when you create the CrustType
	// slot. If you have defined a fallback intent the abort statement will not be sent
	// to the user, the fallback intent is used instead. For more information, see
	// AMAZON.FallbackIntent
	// (https://docs.aws.amazon.com/lex/latest/dg/built-in-intent-fallback.html).
	AbortStatement *types.Statement
	// Identifies a specific revision of the $LATEST version. When you create a new
	// bot, leave the checksum field blank. If you specify a checksum you get a
	// BadRequestException exception. When you want to update a bot, set the checksum
	// field to the checksum of the most recent revision of the $LATEST version. If you
	// don't specify the  checksum field, or if the checksum does not match the $LATEST
	// version, you get a PreconditionFailedException exception.
	Checksum *string
	// The name of the bot. The name is not case sensitive.
	Name *string
	// When Amazon Lex doesn't understand the user's intent, it uses this message to
	// get clarification. To specify how many times Amazon Lex should repeat the
	// clarification prompt, use the maxAttempts field. If Amazon Lex still doesn't
	// understand, it sends the message in the abortStatement field. When you create a
	// clarification prompt, make sure that it suggests the correct response from the
	// user. for example, for a bot that orders pizza and drinks, you might create this
	// clarification prompt: "What would you like to do? You can say 'Order a pizza' or
	// 'Order a drink.'" If you have defined a fallback intent, it will be invoked if
	// the clarification prompt is repeated the number of times defined in the
	// maxAttempts field. For more information, see  AMAZON.FallbackIntent
	// (https://docs.aws.amazon.com/lex/latest/dg/built-in-intent-fallback.html). If
	// you don't define a clarification prompt, at runtime Amazon Lex will return a 400
	// Bad Request exception in three cases:
	//
	//     * Follow-up prompt - When the user
	// responds to a follow-up prompt but does not provide an intent. For example, in
	// response to a follow-up prompt that says "Would you like anything else today?"
	// the user says "Yes." Amazon Lex will return a 400 Bad Request exception because
	// it does not have a clarification prompt to send to the user to get an intent.
	//
	//
	// * Lambda function - When using a Lambda function, you return an ElicitIntent
	// dialog type. Since Amazon Lex does not have a clarification prompt to get an
	// intent from the user, it returns a 400 Bad Request exception.
	//
	//     * PutSession
	// operation - When using the PutSession operation, you send an ElicitIntent dialog
	// type. Since Amazon Lex does not have a clarification prompt to get an intent
	// from the user, it returns a 400 Bad Request exception.
	ClarificationPrompt *types.Prompt
	// An array of Intent objects. Each intent represents a command that a user can
	// express. For example, a pizza ordering bot might support an OrderPizza intent.
	// For more information, see how-it-works ().
	Intents []*types.Intent
	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions
	// with the user. The locale configured for the voice must match the locale of the
	// bot. For more information, see Voices in Amazon Polly
	// (https://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly
	// Developer Guide.
	VoiceId *string
	// A description of the bot.
	Description *string
	// When set to true user utterances are sent to Amazon Comprehend for sentiment
	// analysis. If you don't specify detectSentiment, the default is false.
	DetectSentiment *bool
	// Specifies the target locale for the bot. Any intent used in the bot must be
	// compatible with the locale of the bot.  <p>The default is
	// <code>en-US</code>.</p>
	Locale types.Locale
	// A list of tags to add to the bot. You can only add tags when you create a bot,
	// you can't use the PutBot operation to update the tags on a bot. To update tags,
	// use the TagResource operation.
	Tags []*types.Tag
}

type PutBotOutput struct {
	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user. For more information, see PutBot ().
	VoiceId *string
	// The prompts that Amazon Lex uses when it doesn't understand the user's intent.
	// For more information, see PutBot ().
	ClarificationPrompt *types.Prompt
	// true if the bot is configured to send user utterances to Amazon Comprehend for
	// sentiment analysis. If the detectSentiment field was not specified in the
	// request, the detectSentiment field is false in the response.
	DetectSentiment *bool
	// Checksum of the bot that you created.
	Checksum *string
	// A description of the bot.
	Description *string
	// The target locale for the bot.
	Locale types.Locale
	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to a
	// website, program, or other application that is directed or targeted, in whole or
	// in part, to children under age 13 and subject to COPPA. By specifying false in
	// the childDirected field, you confirm that your use of Amazon Lex is not related
	// to a website, program, or other application that is directed or targeted, in
	// whole or in part, to children under age 13 and subject to COPPA. You may not
	// specify a default value for the childDirected field that does not accurately
	// reflect whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. If your use of Amazon Lex relates to a
	// website, program, or other application that is directed in whole or in part, to
	// children under age 13, you must obtain any required verifiable parental consent
	// under COPPA. For information regarding the use of Amazon Lex in connection with
	// websites, programs, or other applications that are directed or targeted, in
	// whole or in part, to children under age 13, see the Amazon Lex FAQ.
	// (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool
	// A list of tags associated with the bot.
	Tags []*types.Tag
	// If status is FAILED, Amazon Lex provides the reason that it failed to build the
	// bot.
	FailureReason *string
	// An array of Intent objects. For more information, see PutBot ().
	Intents []*types.Intent
	// The version of the bot. For a new bot, the version is always $LATEST.
	Version *string
	// The maximum length of time that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot ().
	IdleSessionTTLInSeconds *int32
	// The name of the bot.
	Name *string
	// The message that Amazon Lex uses to abort a conversation. For more information,
	// see PutBot ().
	AbortStatement *types.Statement
	// When you send a request to create a bot with processBehavior set to BUILD,
	// Amazon Lex sets the status response element to BUILDING. In the
	// READY_BASIC_TESTING state you can test the bot with user inputs that exactly
	// match the utterances configured for the bot's intents and values in the slot
	// types. If Amazon Lex can't build the bot, Amazon Lex sets status to FAILED.
	// Amazon Lex returns the reason for the failure in the failureReason response
	// element. When you set processBehavior to SAVE, Amazon Lex sets the status code
	// to NOT BUILT. When the bot is in the READY state you can test and publish the
	// bot.
	Status types.Status
	// The date that the bot was updated. When you create a resource, the creation date
	// and last updated date are the same.
	LastUpdatedDate *time.Time
	// True if a new version of the bot was created. If the createVersion field was not
	// specified in the request, the createVersion field is set to false in the
	// response.
	CreateVersion *bool
	// The date that the bot was created.
	CreatedDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutBotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutBot{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBot{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "PutBot",
	}
}