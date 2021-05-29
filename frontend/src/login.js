import LoginApp from "./components/LoginApp.svelte";

export const loginApp = new LoginApp({
	target: document.getElementById("svelte-login"),
	props: {
		clientId: window.clientId
	}
});
