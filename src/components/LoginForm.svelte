<script lang="typescript">
  import { blur } from "svelte/transition";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  import UsersTable from "./UsersTable.svelte"

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
  <UsersTable users={users} />
  <br />
  
  {#if loginError}
    <div class="error" transition:blur="{{duration: 500}}">
      <strong>Error</strong> Login information incorrect
    </div>
  {/if}

  <div>
    <form action="#">
      <label for="username">Username</label>
      <input type="text" name="username" placeholder="Username" bind:value={username} class:error="{loginError}">

      <label for="password">Password</label>
      <input type="password" name="password" placeholder="Password" bind:value={password} class:error="{loginError}">

      <button on:click|preventDefault={handleLogin}>Login</button>
    </form> 
  </div>
</section>
