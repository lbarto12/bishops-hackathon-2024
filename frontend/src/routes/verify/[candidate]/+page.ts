import {getCandidateInt} from "$lib/voting/verify";

export async function load({params}) {
    let candidate_display: string = await getCandidateInt(params.candidate);
    console.log(candidate_display);

    return {
        candidate: params.candidate,
        candidate_display: candidate_display,
        candidates: { // :'(
            1: "Candidate 1",
            2: "Candidate 2",
            3: "Candidate 3",
        }
    }
}
