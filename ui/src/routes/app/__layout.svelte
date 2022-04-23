<script lang="ts">
  import { goto } from "$app/navigation";
  import { user } from "../../stores/user";
  import { beforeUpdate } from 'svelte';

  const logout = () => {
    sessionStorage.setItem("token", "");
    user.clear();
    goto("/sign-in");
  }

  beforeUpdate(() => {
    const token = sessionStorage.getItem("token");

    if (!token) {
      goto("/sign-in")
    }
  });

</script>

<div>
  <button on:click={logout}>logout</button>
</div>
<slot></slot>