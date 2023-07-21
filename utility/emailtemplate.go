package utility

import (
	"bytes"
	"html/template"
)

type EmailData struct {
	Name       string
	Email      string
	Link       string
	ExpiryTime string
}

// GenerateHTMLEmail generates the HTML content for the email using the provided data
func GenerateHTMLEmail(data EmailData) (string, error) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Bible Compass Forgot Password</title>
		<style>
		  body {
			font-family: Arial, sans-serif;
			line-height: 1.6;
			margin: 0;
			padding: 0;
			background-color: #f5f5f5;
		  }
		  .container {
			max-width: 600px;
			margin: 0 auto;
			padding: 20px;
			border: 1px solid #ccc;
			border-radius: 5px;
			box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
			background-color: #fff;
		  }
		  .header {
			background-color: #0ba37f;
			padding: 15px;
			text-align: center;
			border-radius: 5px 5px 0 0;
		  }
		  .header h1 {
			color: #fff;
			margin: 0;
			padding: 0;
		  }
		  .content {
			padding: 20px 0;
			color: #333;
		  }
		  .button {
			display: inline-block;
			background-color: #0ba37f;
			color: #fff;
			text-decoration: none;
			padding: 10px 20px;
			border-radius: 5px;
			margin-top: 10px;
		  }
		  .footer {
			text-align: center;
			padding-top: 20px;
			color: #777;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <div class="header">
			<h1>Forgot Password Email</h1>
		  </div>
		  <div class="content">
			<p>Hello {{.Email}},</p>
			<p>
			  Welcome to our platform! To complete your verification, please verify
			  your email address by clicking the button below and proceed to choose
			  a new password:
			</p>
			<a href="{{.Link}}" class="button">Verify Email</a>
			<p>
			  If the button above does not work, you can also copy and paste the
			  following link into your browser:
			</p>
			<p>{{.Link}}</p>
			<p>This verification link will expire in {{.ExpiryTime}}.</p>
			<p>
			  If you did not initiate a forgot password process on our platform,
			  please disregard this email.
			</p>
			<p>Thank you,<br />The Bible Compass Team</p>
		  </div>
		  <div class="footer">
			<p>This is an automated email, please do not reply.</p>
		  </div>
		</div>
	  </body>
	</html>`

	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var emailContent bytes.Buffer
	err = tmpl.Execute(&emailContent, data)
	if err != nil {
		return "", err
	}

	return emailContent.String(), nil
}

// GenerateHTMLEmail generates the HTML content for the email using the provided data
func GenerateRegistrationHTMLEmail(data EmailData) (string, error) {
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Congratulations on Your Registration!</title>
		<style>
		  body {
			font-family: Arial, sans-serif;
			line-height: 1.6;
			margin: 0;
			padding: 0;
			background-color: #f5f5f5;
		  }
		  .container {
			max-width: 600px;
			margin: 0 auto;
			padding: 20px;
			border: 1px solid #ccc;
			border-radius: 5px;
			box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
			background-color: #fff;
		  }
		  .header {
			background-color: #0ba37f;
			padding: 15px;
			text-align: center;
			border-radius: 5px 5px 0 0;
		  }
		  .header h1 {
			color: #fff;
			margin: 0;
			padding: 0;
		  }
		  .content {
			padding: 20px 0;
			color: #333;
		  }
	
		  .footer {
			text-align: center;
			padding-top: 20px;
			color: #777;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <div class="header">
			<h1>Congratulations on Your Registration!</h1>
		  </div>
		  <div class="content">
			<p>Hello {{.Email}},</p>
			<p>
			  Congratulations! You have successfully registered on our platform. We
			  are excited to have you as part of our community.
			</p>
			<p>
			  With your new account, you can access exclusive features, interact
			  with other members, and enjoy our premium services.
			</p>
			<p>
			  If you have any questions or need assistance, please don't hesitate to
			  reach out to our support team.
			</p>
			<p>Once again, congratulations and welcome aboard!</p>
			<p>Best regards,<br />The Bible Compass Team</p>
		  </div>
		  <div class="footer">
			<p>This is an automated email, please do not reply.</p>
		  </div>
		</div>
	  </body>
	</html>	   
`

	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var emailContent bytes.Buffer
	err = tmpl.Execute(&emailContent, data)
	if err != nil {
		return "", err
	}

	return emailContent.String(), nil
}
