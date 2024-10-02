
<script>
    import Chart from "chart.js/auto";
    let chartObject;
    let dataToGraph = [0, 0, 0];  // Initialize with 0, will be updated by API
    let errorVoteCount = null;

    // Define the colors for each candidate
    const candidateColors = {
        candidate1: {
            backgroundColor: "rgba(255, 99, 132, 0.7)",  // Color for Candidate 1
            borderColor: "rgba(255, 99, 132, 1)"         // Border color for Candidate 1
        },
        candidate2: {
            backgroundColor: "rgba(54, 162, 235, 0.7)",  // Color for Candidate 2
            borderColor: "rgba(54, 162, 235, 1)"         // Border color for Candidate 2
        },
        candidate3: {
            backgroundColor: "rgba(75, 192, 192, 0.7)",  // Color for Candidate 3
            borderColor: "rgba(75, 192, 192, 1)"         // Border color for Candidate 3
        }
    };

    // Function to calculate percentages
    function calculatePercentages(data) {
        const total = data.reduce((sum, value) => sum + value, 0);
        if (total === 0) {
            return [0, 0, 0];  // Return 0% for all candidates if total is 0
        }
        return data.map((value) => ((value / total) * 100).toFixed(2));
    }

    // Fetch vote data from the API
    async function fetchVoteData() {
        try {
            const res = await fetch("http://localhost:5006/api/vote-counts");
            if (!res.ok) {
                throw new Error(`Failed to fetch vote data, status: ${res.status}`);
            }
            const responseData = await res.json();

            // Extract the 'votes' from the 'data' array
            if (responseData && Array.isArray(responseData.data)) {
                const votes = responseData.data.map(candidate => candidate.votes);
                return votes;
            } else {
                throw new Error("Invalid data format: 'data' is missing or not an array");
            }
        } catch (err) {
            console.error("Error fetching vote data:", err);
            errorVoteCount = err.message;
            return [0, 0, 0];  // Return default values if fetch fails
        }
    }

    // Function to update the chart with new data from the API
    async function updateChartWithAPIData() {
        const fetchedData = await fetchVoteData();
        if (fetchedData.length > 0) {
            const percentages = calculatePercentages(fetchedData);
            chartObject.data.datasets[0].data = percentages;  // Update chart data
            chartObject.update();  // Redraw the chart with new data
        }
    }

    // Set up the chart
    function chart(node, data) {
        chartObject = new Chart(node, {
            type: "pie",  // Pie chart
            data: {
                labels: ["Candidate 1", "Candidate 2", "Candidate 3"],
                datasets: [
                    {
                        label: "#Vote Percentage",
                        data: calculatePercentages(data),  // Initialize chart with data
                        backgroundColor: [
                            candidateColors.candidate1.backgroundColor,
                            candidateColors.candidate2.backgroundColor,
                            candidateColors.candidate3.backgroundColor,
                        ],
                        borderColor: [
                            candidateColors.candidate1.borderColor,
                            candidateColors.candidate2.borderColor,
                            candidateColors.candidate3.borderColor,
                        ],
                        borderWidth: 1,
                    },
                ],
            },
            options: {
                plugins: {
                    tooltip: {
                        callbacks: {
                            label: function(context) {
                                return context.raw + "%";  // Show percentage in tooltip
                            }
                        }
                    }
                }
            },
        });
    }

    // When the component is mounted, set up the chart and start polling
    import { onMount } from "svelte";
    onMount(async () => {
        // Immediately fetch and update the chart with the initial data
        const initialData = await fetchVoteData();
        if (initialData.length > 0) {
            dataToGraph = initialData;
            const percentages = calculatePercentages(initialData);
            chartObject.data.datasets[0].data = percentages;
            chartObject.update();
        }

        // Poll the API every 5 seconds to check for new data
        setInterval(async () => {
            await updateChartWithAPIData();  // Fetch and update chart every 5 seconds
        }, 5000);
    });
</script>

<!-- HTML Structure -->
<div class="container">
    <canvas class="chart" use:chart={dataToGraph}></canvas>
    <h1><strong>Pie Chart</strong></h1>
</div>

<style>
    .container {
        position: absolute;
        top: 0;
        right: 0;
        width: auto;
        height: auto;
        margin: 20px; /* Add more margin */
    }

    .chart {
        width: 400px;
        height: 300px;
    }
</style>

