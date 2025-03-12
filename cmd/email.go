package cmd

import (
	"fmt"
	"log"
	"os"

	gomail "gopkg.in/mail.v2"
)

func SendEmail() {
	username := "studywithmelimited@gmail.com"
	token := os.Getenv("APPKEYCANVAS")

	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("To", "davematteer@gmail.com",
		//"Wumpinilatif7779@outlook.com",
		//"neamewu@gmail.com",
		//"annabelkudiabor70@gmail.com",
		//"charlesjesimiel@gmail.com",
		//"curtis.adzornu@gmail.com",
		//"gwumah@gmail.com",
		//"nksarpongs15@gmail.com",
		"maameossei@gmail.com",
		"Addokufuor.ka@gmail.com",
	)
	message.SetHeader("Subject", "Welcome to StudyWithMe - Your Personalized Learning tool!!!")

	body := `
		<html>
		<head>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					background: #ffffff;
					padding: 20px;
					border-radius: 10px;
					box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
				}
				h2 {
					color: #2c3e50;
					text-align: center;
				}
				p {
					color: #555555;
					font-size: 16px;
					line-height: 1.6;
				}
				.footer {
					text-align: center;
					font-size: 14px;
					color: #888888;
					margin-top: 20px;
				}
				.button {
					display: inline-block;
					background-color: #3498db;
					color: #ffffff;
					padding: 10px 20px;
					text-decoration: none;
					border-radius: 5px;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h2>🎉 Thank You for Using Our Services! 🎉</h2>
				<p>Hello,</p>
				<p>We appreciate you using our platform. Attached is your <strong>personalized study plan</strong> in PDF format.</p>
				<p>We hope this helps you stay on track with your studies!</p>
				<p style="text-align: center;">
					<a href="#" class="button">View Your Study Plan</a>
				</p>
				<div class="footer">
					<p>Best regards,<br><strong>Study With Me Limited</strong></p>
				</div>
			</div>
		</body>
		</html>
	`
	message.SetBody("text/html", body)
	message.Attach("study_plan.pdf", gomail.SetHeader(map[string][]string{
		"Content-Type": {"application/pdf"},
	}))

	dialer := gomail.NewDialer("smtp.gmail.com", 465, username, token)

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully!!!")
}
