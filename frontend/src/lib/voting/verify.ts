
import {PUBLIC_API_HOST} from "$env/static/public";
import {sha256} from "js-sha256";


export async function vote(name: string, healthCard: string, candidate: string): Promise<Response> {

    const hash = sha256.create();
    hash.update(`${name}${healthCard}${candidate}`);
    const final: string = hash.hex();

    return await fetch(PUBLIC_API_HOST + "/vote", {
        method: "POST",
        body: JSON.stringify({
            data: final
        })
    });
}
