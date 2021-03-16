<script lang="typescript">
  import { blur } from "svelte/transition";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  let errorMessage = "";
  let name = "";
  let redirect_uri = "";
  let setupError = false;

  function successfulSetup(data: any) {
    dispatch("success", {
			data: data,
      name: name,
      redirect_uri: redirect_uri
		});
  }

  function error(msg: string) {
    setupError = true;
    errorMessage = msg;
    setTimeout(() => setupError = false, 5000)
  }

  function getCodes() {
    const url = "http://localhost:8090/client"
    const req = { name: name, redirect_uri: redirect_uri };
    const config = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      body: JSON.stringify(req)
    }
    
    fetch(url, config)
    .then(async r => await r.json())
    .then(function(data) {
      console.log("data");
      console.log(data);
      successfulSetup(data);
    }).catch(err => console.log('Request Failed', err));
  }

  function handleSubmit() {
    if (name.length == 0 || redirect_uri.length == 0) {
      error("Setup information missing");
    } else if ((redirect_uri.match(/[&\?<>{}\[\]=]/)) != null) {
      error("Redirect URI contains an invalid character");
    } else {
      getCodes();
    }
  }
</script>

<style>
  div.error {
    background-color: rgba(255, 0, 0, 0.3);
    border: 2px solid #990000;
    border-radius: 6px;
    margin-bottom: 12px;
    padding: 12px;
    width: 100%;
  }

  input.error {
    border: 2px solid #990000;
  }
</style>

<div>
  {#if setupError}
    <div class="error" transition:blur="{{duration: 500}}">
      <strong>Error</strong> {errorMessage}
    </div>
  {/if}

  <form action="#">
    <label for="name">Application Name</label>
    <input type="text" name="name" placeholder="Unique Name" bind:value={name} class:error="{setupError}">

    <label for="redirect_uri">Redirect URI</label>
    <input type="text" name="redirect_uri" placeholder="https://" bind:value={redirect_uri} class:error="{setupError}">

    <br /><button on:click|preventDefault={handleSubmit}>Generate</button>
  </form> 
</div>