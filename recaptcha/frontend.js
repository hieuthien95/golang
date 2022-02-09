import { Button, FormLabel, TextField, Typography } from "@material-ui/core"
import { GoogleReCaptchaProvider, useGoogleReCaptcha } from "react-google-recaptcha-v3"

export async function getServerSideProps(ctx) {
	return { props: {} }
}

const RegisterPage = (props) => {
	const { executeRecaptcha } = useGoogleReCaptcha()

	const onSubmitRegister = async (formData) => {
		try {
			const newToken = await executeRecaptcha("register")
			console.log(newToken)
		} catch (err) {
			throw new Error("Token error")
		}

		// do action
	}

	return (
		<form>
			<Typography gutterBottom>
				<FormLabel>Tên công ty</FormLabel>
			</Typography>

			<TextField fullWidth variant="outlined" size="small" type="text" placeholder="Nhập tên nhà cung cấp" />

			<Button variant="contained" color="primary" onClick={onSubmitRegister}>
				Đăng ký bán hàng
			</Button>
		</form>
	)
}
const AppCapt = () => {
	return (
		<GoogleReCaptchaProvider language="en" reCaptchaKey="<site_key>">
			<RegisterPage />
		</GoogleReCaptchaProvider>
	)
}

export default AppCapt
