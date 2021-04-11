<script lang="typescript">
  export let users: { username: string, password: string }[]
</script>

<div>
  <h2>Provided Users</h2>
  <p>
    We have provided you with three unique users to help you test your application. Each has their
    username, password and unique data enumerated below for your convenience.
  </p>
  <table>
    <thead>
      <tr>
        <th>Username</th>
        <th>Password</th>
      </tr>
    </thead>
    <tbody>
      {#each users as { username, password }}
        <tr>
          <td>{ username }</td>
          <td>{ password }</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
