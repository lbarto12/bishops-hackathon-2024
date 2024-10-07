import {writable} from 'svelte/store'; //
import {PUBLIC_API_HOST} from "$env/static/public";

// To store to hold poll data
export const pollsData = writable([]);


export async function fetchPollData() {
    const url = PUBLIC_API_HOST + "/tabulation/polls";
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error('Network response was not ok as it should');
        }
        const data = await response.json();

        // To get the total votes
        const totalVotes = data.data.polls.reduce((sum, poll) => sum + poll.votes, 0);

        // To percentage calculation
        const pollsWithPercentages = data.data.polls.map(poll => ({
            ...poll, // Copy all the properties from the POLL object
            percentage: ((poll.votes / totalVotes) * 100).toFixed(2)
        }));

        pollsData.set(pollsWithPercentages);  // Update the store with new data we got
        return pollsWithPercentages;  // Return the mapped data
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        return [];
    }
}

export function startPolling(interval = 5000) {
    fetchPollData(); // Fetch immediately when polling starts

    const pollingInterval = setInterval(() => {
        fetchPollData(); // Fetch the data every `interval` milliseconds
    }, interval);

    // Return a function to stop polling (use this for cleanup)
    return () => clearInterval(pollingInterval);
}
