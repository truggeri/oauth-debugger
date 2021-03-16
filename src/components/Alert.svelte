<script lang="typescript">
  import { blur } from "svelte/transition";

  export let klass = "";
  export let boldMsg = "";
  export let message = "";
</script>

<style>
  div.alert {
    border-radius: 6px;
    margin-bottom: 12px;
    padding: 12px;
    width: 100%;
  }

  div.error {
    background-color: rgba(255, 0, 0, 0.3);
    border: 2px solid #990000;
  }

  div.success {
    background-color: rgba(0, 255, 0, 0.3);
    border: 2px solid #009900;
  }
</style>

<div class="alert {klass}" transition:blur="{{duration: 500}}">
  {#if boldMsg != ""}<strong>{ boldMsg }</strong>{/if} { message }
</div>