
export function load({params}) {
    return {
        candidate: params.candidate,
        candidates: { // :'(
            1: "Candidate 1",
            2: "Candidate 2",
            3: "Candidate 3",
        }
    }
}
