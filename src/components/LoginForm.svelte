<script lang="typescript">
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();

  import UsersTable from "./UsersTable.svelte"

  export let users: {username: string, password: string}[]

  let username: string
  let password: string

  function matchUser(user: { username: string, password: string }): boolean {
    return user.username == username && user.password == password
  }

  function handleLogin() {
    if (users.map(matchUser).some(b => b)) {
      dispatch("login", {
        username: username
      });
    }
	}
</script>

<section>
  <UsersTable users={users} />
  <br />

  <div>
    <form action="#">
      <label for="username">Username</label>
      <input type="text" name="username" placeholder="Username" bind:value={username}>

      <label for="password">Password</label>
      <input type="password" name="password" placeholder="Password" bind:value={password}>

      <button on:click|preventDefault={handleLogin}>Login</button>
    </form> 
  </div>
</section>
