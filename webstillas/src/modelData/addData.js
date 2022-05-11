import React from 'react'
import fetchModel from "./fetchData";
import { useQuery } from 'react-query'

//Todo set timeout
export const GetDummyData = (dataName, url) => {
    const { isLoading, data, isError, isLoadingError} = useQuery(dataName, ()=>{
        return fetchModel(url)
    }, {
        refetchOnMount: false,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false

    })

    return {isLoading, data, isError, isLoadingError}
}



