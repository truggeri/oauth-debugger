<script lang="typescript">
    import { blur } from "svelte/transition";

    let name = "";
    let redirect_uri = "";
    let setupError = false;
    let errorMsg = "";

    function error(msg: string) {
      setupError = true;
        errorMsg = msg;
        setTimeout(() => setupError = false, 5000)
    }

    async function getCodes() {
      const url = "/client"
      const response = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: name, redirect_uri: redirect_uri })
      });
      return response.json();
    }

    function handleSubmit() {
      if (name.length == 0 || redirect_uri.length == 0) {
        error("Setup information missing");
      } else if ((redirect_uri.match(/[&\?<>{}\[\]=]/)) != null) {
        error("Redirect URI contains an invalid character");
      } else {
        getCodes().then(data => {
          console.log(data);
        });
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

<section>
  {#if setupError}
    <div class="error" transition:blur="{{duration: 500}}">
      <strong>Error</strong> {errorMsg}
    </div>
  {/if}

  <div>
    <form action="#">
      <label for="name">Application Name</label>
      <input type="text" name="name" placeholder="Unique Name" bind:value={name} class:error="{setupError}">

      <label for="redirect_uri">Redirect URI</label>
      <input type="text" name="redirect_uri" placeholder="https://" bind:value={redirect_uri} class:error="{setupError}">

      <button on:click|preventDefault={handleSubmit}>Generate</button>
    </form> 
  </div>
</section>