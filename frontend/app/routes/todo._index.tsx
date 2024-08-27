import {json, useLoaderData} from "@remix-run/react";

export const loader = async () => {
    const response = await fetch(`http://localhost:8080/todos`,
    {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "GET",
    });

    const data = await response.json();
    
    return json({data});
}


export default function TodoIndex() {
    const {data} = useLoaderData<typeof loader>();

    return (
        <div>
            <nav>
                {data.length ? (
                    <ul>
                        {data.map((todo) => (
                            <li key={todo.id}>
                                <h2>{todo.title}</h2>                      
                            </li>
                        ))}
                    </ul>
                ) : (
                    <p>
                        <i>No todo list</i>
                    </p>
                )}
            </nav>
        </div>
    )
}
