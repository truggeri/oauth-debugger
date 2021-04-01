import LoginApp from "./components/LoginApp.svelte";

export const loginApp = new LoginApp({
	target: document.getElementById("login-app"),
	props: {
		clientId: window.clientId
	}
});
