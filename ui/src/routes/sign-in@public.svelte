<script lang="ts">
  let email = "";
  let password = "";

  const handleSubmit = (e) => {
    console.log("submit event", e);
  };
</script>

<div class="flex justify-center h-screen items-center">
  <div class="card w-96 bg-base-100 bg-slate-400 shadow-xl h-1/3 text-black nav-offset card-min-height">
    <div class="card-body justify-between">
      <div>
        <h1 class="card-title pb-5">Sign In</h1>
        <form id="sign-in" on:submit|preventDefault={handleSubmit}>
          <div class="form-control w-full max-w-xs">
            <label for="email" class="label">
              <span class="label-text-alt text-black">Email</span>
            </label>
            <input bind:value={email} name="email" type="text" placeholder="Email" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">Error</span>
          </div>
          <div class="form-control w-full max-w-xs">
            <label for="password" class="label">
              <span class="label-text-alt text-black">Password</span>
            </label>
            <input bind:value={password} name="password" type="password" placeholder="Password" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">Error</span>
          </div>
        </form>
      </div>
      <div class="card-actions">
        <button class="btn btn-primary btn-block">Sign In</button>
      </div>
    </div>
  </div>
</div>

<style>
  .nav-offset {
    margin-bottom: 30vh;
  }

  .input-error {
    min-height: 16px;
  }

  .card-min-height {
    min-height: 400px;
  }
</style>