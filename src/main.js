import LoginApp from "./components/LoginApp.svelte";
import SetupApp from "./components/SetupApp.svelte";

export const loginApp = new LoginApp({
	target: document.getElementById("login-app"),
	props: {
		clientId: window.clientId
	}
});

export const setupApp = new SetupApp({
	target: document.getElementById("setup-app"),
	props: {}
});
