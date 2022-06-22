import type { InferGetStaticPropsType, NextPage } from 'next'

import Image from 'next/image'

type Task = {
    id: string,
    title: string,
    description: string,
    duedate: string,
    priority: string,
  }

export const getStaticProps = async () => {
    const res = await fetch("http://localhost:3000/api/tasks")
    const data: Task[] = await res.json()
  
    return {
      props: {
        data,
      },
    }
  }

const TestFrom = ({ data }: InferGetStaticPropsType<typeof getStaticProps>) => {
    let array: any[] = []
    for (let i = 0; i < data.length; i++)
      array.push(
        <a href="#" className="text-gray-700 block px-4 py-2 text-sm" role="menuitem" id="menu-item-0">{data[i].id} - {data[i].title}</a>
    )

    const handleSubmit = async (event: any) => {
        event.preventDefault()

        const data = {
            title: event.target.title.value,
        }
        const JSONdata = JSON.stringify(data)
        const response = await fetch('http://localhost:3000/api/tasks', {
            // Body of the request is the JSON data we created above.
            body: JSONdata,
      
            // Tell the server we're sending JSON.
            headers: {
              'Content-Type': 'application/json',
            },
            // The method is POST because we are sending data.
            method: 'POST',
        })
        const result = await response.json()
    }
    


  return (
    <>
    <div className="container">
      <form onSubmit={handleSubmit}>
      <div className="shadow sm:rounded-md sm:overflow-hidden">
            <div className="px-4 py-5 bg-white space-y-6 sm:p-6">
              <div className="grid grid-cols-3 gap-6">
                <div className="col-span-3 sm:col-span-2">
        <label htmlFor="title" className="block text-sm font-medium text-gray-700">Title:</label>
        <div className="mt-1 flex rounded-md shadow-sm">
        <input className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md" type="text" id="title" name="title" />
        </div>
        </div>
        </div>
        </div>
        </div>
        <button type="submit">Submit</button>
      </form>
       {array}
    </div>
    
    </>
  )
}

export default TestFrom


