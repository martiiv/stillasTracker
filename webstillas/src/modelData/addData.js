import React from 'react'
import fetchModel from "./fetchData";
import { useQuery } from 'react-query'

//Todo set timeout

/**
 * Function that will fetch data from api, and cache data
 *
 * @param dataName key to data caching
 * @param url to the api
 * @returns {{isLoading: boolean, isLoadingError: boolean, isError: boolean, data: unknown}}
 */
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



