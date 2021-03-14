<script lang="typescript">
  import { blur, slide } from 'svelte/transition';
  import LoginForm from "./LoginForm.svelte"
  import AuthorizeForm from "./AuthorizeForm.svelte"
  import users from "./users"

  let showAuthorize = false
  let user: {username: string, password: string}

  function handleLogin(e: any) {
    showAuthorize = true
    user = e.detail.user
  }
</script>

<section>
  {#if showAuthorize}
    <div out:blur="{{duration: 500}}" in:slide="{{delay: 500, duration: 500}}">
      <AuthorizeForm on:deny={() => showAuthorize = false} user={user} />
    </div>
  {:else}
    <div out:blur="{{duration: 500}}" in:slide="{{delay: 500, duration: 500}}">
      <LoginForm users={users} on:login={handleLogin} />
    </div>
  {/if}
</section>