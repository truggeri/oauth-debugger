<script lang="typescript">
  import SetupForm from "./SetupForm.svelte"
  import SetupSuccess from "./SetupSuccess.svelte"

  let client: any;
  let success = false;

  function handleSuccess(event: any) {
    client = event.detail;
    success = true;
  }
</script>

<section>
  {#if success}
    <SetupSuccess client={client} />
  {:else}
    <SetupForm on:success={handleSuccess} />
  {/if}
</section>