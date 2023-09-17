export async function PingServer(): Promise<{statusCode: number, message: string}> {
    const response = await fetch("http://localhost:8080/api/ping");
    let statusCode =  response.status
    let responseJSON = await response.json()
    return {statusCode, message: responseJSON["message"]}
}