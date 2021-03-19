<script lang="typescript">
  import Alert from "./Alert.svelte";
  import SetupForm from "./SetupForm.svelte"
  import SetupSuccess from "./SetupSuccess.svelte"

  let client: any;
  let error = false;
  let success = false;

  function handleSuccess(data: any) {
    client = data;
    success = true;
  }

  function getClient(event: any) {
    const url = CONFIG.default.app_url + "/client"
    const req = { name: event.detail.name, redirect_uri: event.detail.redirect_uri };
    const config = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      body: JSON.stringify(req)
    }
    
    fetch(url, config)
    .then(async function(resp) {
      if (resp.ok) {
        return handleSuccess(await resp.json());
      }
      error = true;
    })
    .catch(function(err) {
      console.log("Request Failed", err);
      error = true;
    });
  }
</script>

<section>
  {#if error}
    <Alert klass="error" boldMsg="Error" message="Something went wrong" />
  {/if}

  {#if success}
    <SetupSuccess client={client} />
  {:else}
    <SetupForm on:submit={getClient} />
  {/if}
</section>