
// --- 1. Erweiterte Typ-Definitionen ---

import type GPX from "$lib/models/gpx/gpx";
import type Track from "$lib/models/gpx/track";
import PocketBase from 'pocketbase';
import { env } from "process";


/** Helper: return first element if array */
function first<T>(v: T | T[] | undefined): T | undefined {
    if (v === undefined) return undefined;
    return Array.isArray(v) ? v[0] : v;
}

//** Find a string value for `element` inside possibly nested extension structures */
function findExtensionValue(obj: object, element: string): string | null {
    if (!obj || typeof obj !== 'object') return null;
    for (const [k, v] of Object.entries(obj)) {
        if (k === element || k.endsWith(':' + element) || k.includes(element)) {
            const candidate = first(Array.isArray(v) ? v : [v]);
            if (typeof candidate === 'string')
                return candidate;
        }
    }
    return null;
}

export async function getCategoryFromGPX(gpx: GPX): Promise<string | null> {
    if (!gpx.trk) return null

    const track: Track = Array.isArray(gpx.trk) ? gpx.trk[0] : gpx.trk;

    // 1. Metadaten Check (Tags & Extensions)
    // Hinweis: Hier geben wir generische Typen zurück, die wir später evtl. verfeinern
    const tagResult = await checkTagsAndExtensions(gpx, track);

    return tagResult;
}

async function mapCategoryToActivity(categoryString: string): Promise<string | null> {
    const target = categoryString.toLowerCase();
    const pb = new PocketBase(env.PUBLIC_POCKETBASE_URL);
    const record = await pb.collection('categories_aliases').getFirstListItem('alias="' + target + '"', 
        { expand: 'category' });

    
    // Hm, wie komme ich jetzt an die Category-Definitionen und deren Aliase in db?
    const result = record.id;
    return result
}

async function mapCategoryJsonToActivity(categoryJson: object, element: string): Promise<string | null> {
    const val = findExtensionValue(categoryJson, element);
    if (!val || typeof val !== 'string') return null;
    return await mapCategoryToActivity(val);
}

async function checkTagsAndExtensions(gpx: GPX, track: Track): Promise<string | null> {
    // A. Das Standard <type> Tag (auch von OsmAnd genutzt)
    if (track.type) {
        const t = track.type.toLowerCase();
        return await mapCategoryToActivity(t)
    }

    // B. Extensions prüfen (OsmAnd & Garmin)
    if (track.extensions) {
        // { 'topografix:color': 'c0c0c0' }
        console.log('Track Extensions:', track.extensions);

        return mapCategoryJsonToActivity(track.extensions, 'activity_type')
    }

    // B. Extensions prüfen (OsmAnd & Garmin)
    if (gpx.metadata?.extensions) {
        // { 'osmand:activity_type': 'hiking' }
        console.log('GPX Metadata Extensions:', gpx.metadata.extensions);
        return mapCategoryJsonToActivity(gpx.metadata.extensions, 'osmand:activity')
    }

    return null;
}
