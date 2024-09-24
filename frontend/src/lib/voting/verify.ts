
import {PUBLIC_API_HOST} from "$env/static/public";


export async function vote(name: string, healthCard: string, candidate: string) {
    try{
        await fetch(PUBLIC_API_HOST + "/vote", {
            method: "POST",
            body: JSON.stringify({
                voter: {
                    name: name,
                    health_card: healthCard,
                },
                candidate: candidate,
            })
        });
    }
    catch (e: Error) {
        console.log(e);
    }

    console.log("done");
}
