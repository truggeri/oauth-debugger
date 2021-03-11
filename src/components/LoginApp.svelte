<script lang="typescript">
  import LoginForm from "./LoginForm.svelte"
  import AuthorizeForm from "./AuthorizeForm.svelte"
  import users from "./users"

  let showAuthorize = true

  function handleLogin() {
    showAuthorize = true
  }
</script>

<section>
  {#if showAuthorize}
    <AuthorizeForm user={users[0]} />
  {:else}
    <LoginForm users={users} on:login={handleLogin} />
  {/if}
</section>