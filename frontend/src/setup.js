import SetupApp from "./components/SetupApp.svelte";

export const setupApp = new SetupApp({
	target: document.getElementById("svelte-setup"),
	props: {}
});
