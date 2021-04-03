<script lang="typescript">
  import { blur, slide } from "svelte/transition";
  import Alert from "./Alert.svelte"
  import LoginForm from "./LoginForm.svelte"
  import AuthorizeForm from "./AuthorizeForm.svelte"
  import users from "./users"

  export let clientId: string;

  function getCookieValue(name: string) {
    return document.cookie.match('(^|;)\\s*' + name + '\\s*=\\s*([^;]+)')?.pop() || '';
  }

  let error = false;
  let errorMsg = "";
  let showAuthorize = false
  let user: { username: string, password: string }

  function handleError(msg: string) {
    errorMsg = msg;
    error = true;
    setTimeout(() => error = false, 5000)
  }

  function handleLogin(e: any) {
    showAuthorize = true
    user = e.detail.user
  }

  function handleSuccess(data: any) {
    let uri = data.redirect_uri;
    if (data.success && uri) {
      window.location.replace(uri);
    } else {
      handleError("Something is wrong with the redirect uri provided");
    }
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
        "Accept": "application/json",
        "Content-Type": "application/json",
        "X-Csrf-Token": getCookieValue("__HOST-token"),
      },
      body: JSON.stringify(req),
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
      handleError("Something went wrong on our end")
    });
  }
</script>

<section>
  {#if error}
    <Alert klass="error" boldMsg="Error" message={errorMsg} />
  {/if}

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