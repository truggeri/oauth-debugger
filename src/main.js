import LoginApp from "./components/LoginApp.svelte";

const loginApp = new LoginApp({
	target: document.getElementById("login-app"),
	props: {}
});

export default loginApp;