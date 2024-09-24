<script lang="ts">
    import { vote } from "$lib/voting/verify";
    import {goto} from "$app/navigation";

    export let data;

    console.log(data);

    let legalName: string = "";
    let healthCardNumber: string = "";

    let requesting: boolean = false;
    let failed: boolean = false;

    async function executeVote() {
        if (requesting) return;
        requesting = true;
        const response = await vote(legalName, healthCardNumber, data.candidate);
        console.log(response);
        if (response.status === 201) { // created new
            setTimeout(() => goto('/confirmation'), 0); // stupid
        }
        else if (response.status === 200) { // already voted
            setTimeout(() => goto('/reconfirmation'), 0); // stupid as well
        }
        else {
            failed = true;
            setTimeout(() => {failed = false;}, 5000);
        }
        requesting = false;
    }

</script>

<form class="flex justify-center h-full" on:submit={executeVote}>
    <div class="card w-96 bg-base-100 shadow-xl mt-20 mb-20">
        <div class="card-body">
            <h2 class="card-title">Please Verify Your Identity</h2>
            <div class="items-center mt-2">
                <label class="input input-bordered flex items-center gap-2 mb-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 17c0-2 4-3.1 6-3.1s6 1.1 6 3.1v1H6m9-9a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3a3 3 0 0 1 3 3M3 5v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2"/></svg>
                    <input type="text" class="grow" placeholder="Legal Name" bind:value={legalName} />
                </label>
                <label class="input input-bordered flex items-center gap-2 mb-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 16 16"><path fill="currentColor" d="M15 3v10H1V3zm1-1H0v12h16z"/><path fill="currentColor" d="M9 5h5v1H9zm0 2h5v1H9zm0 2h2v1H9zM6.5 5c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L5 10l2.5-2.1C8.6 6.9 8 5 6.5 5"/></svg>
                    <input type="text" class="grow" placeholder="Health Card Number" bind:value={healthCardNumber} />
                </label>
            </div>
            <div class="card-actions justify-end">

                <button
                        class="btn btn-primary w-full"
                        type="submit">
                    {#if requesting}
                        <span class="loading loading-spinner loading-md"></span>
                    {:else}
                        Submit Vote
                    {/if}
                </button>
            </div>
        </div>
    </div>
</form>

{#if failed}
    <div class="toast toast-center">
        <div class="alert alert-error">
            Something went wrong...
        </div>
    </div>
{/if}



