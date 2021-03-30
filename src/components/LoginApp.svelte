<script lang="typescript">
  import { blur, slide } from "svelte/transition";
  import LoginForm from "./LoginForm.svelte"
  import AuthorizeForm from "./AuthorizeForm.svelte"
  import users from "./users"

  export let clientId: string;

  let error = false;
  let showAuthorize = false
  let user: { username: string, password: string }

  function handleLogin(e: any) {
    showAuthorize = true
    user = e.detail.user
  }

  function grantCode(_event: any) {
    const url = `${CONFIG.default.app_url}/oauth/grant`
    const req = {
      client_id: clientId,
      username: user.username,
    };
    const config: RequestInit = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      },
      body: JSON.stringify(req),
      redirect: "follow",
    }
    
    fetch(url, config)
    .then(function(resp) {
      if (resp.redirected) {
        window.location.replace(resp.url);
      }
    })
    .catch(function(err) {
      console.log("Request Failed", err);
      error = true;
    });
  }
</script>

<section>
  {#if showAuthorize}
    <div out:blur="{{duration: 500}}" in:slide="{{delay: 500, duration: 500}}">
      <AuthorizeForm user={user} on:submit={grantCode} on:deny={() => showAuthorize = false} />
    </div>
  {:else}
    <div out:blur="{{duration: 500}}" in:slide="{{delay: 500, duration: 500}}">
      <LoginForm users={users} on:login={handleLogin} />
    </div>
  {/if}
</section>