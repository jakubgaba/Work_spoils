//NOTE: cookies are not forcefully hard stored
import { redirect } from '@sveltejs/kit';

export const load = async () => {
  return redirect(303, '/mainPage');
};
