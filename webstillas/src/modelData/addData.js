import React from 'react'
import fetchModel from "./fetchData";
import { useQuery } from 'react-query'


export const GetDummyData = (dataName, url) => {
    const { isLoading, data} = useQuery(dataName, ()=>{
        return fetchModel(url)
    }, {
        refetchOnMount: false,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false
    })

    return {isLoading, data}
}


