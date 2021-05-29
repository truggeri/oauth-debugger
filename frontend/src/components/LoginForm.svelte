<script lang="typescript">
  import Alert from "./Alert.svelte";
  import UsersTable from "./UsersTable.svelte"
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let users: {username: string, password: string}[]

  let username: string
  let password: string
  let loginError = false

  function matchUser(user: { username: string, password: string }): boolean {
    return user.username == username && user.password == password
  }

  function handleLogin() {
    if (users.map(matchUser).some(b => b)) {
      dispatch("login", {
        user: { username: username, password: password }
      });
    } else {
      loginError = true
      setTimeout(() => loginError = false, 5000)
    }
	}
</script>

<style>
  input.error {
    border: 2px solid #990000;
  }
</style>

<section>
  <UsersTable users={users} />
  <br />
  
  {#if loginError}
    <Alert klass="error" boldMsg="Error" message="Login information incorrect" />
  {/if}

  <div>
    <h2>Sign in</h2>
    <p>Please sign in to continue.</p>
    <form action="#">
      <label for="username">Username</label>
      <input type="text" name="username" placeholder="Username" bind:value={username} class:error="{loginError}">

      <label for="password">Password</label>
      <input type="password" name="password" placeholder="Password" bind:value={password} class:error="{loginError}">

      <br />
      <button on:click|preventDefault={handleLogin}>Login</button>
    </form> 
  </div>
</section>
