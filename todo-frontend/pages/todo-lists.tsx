import type { NextPage } from 'next'
import { InferGetStaticPropsType } from 'next'

export const getStaticProps = async () => {
  const res = await fetch("http://localhost:3000/api/todo-lists")
  const data: any[] = await res.json()

  return {
      props: {
        data,
      },
  } 
}

const TodoLists = ({ data }: InferGetStaticPropsType<typeof getStaticProps>) => {
  return (
      <>
      <h2>TodoLists</h2>
      {JSON.stringify([data])}
      </>
  );
}

export default TodoLists
