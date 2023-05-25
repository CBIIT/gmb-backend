package gov.nih.nci.bento_ri.model;

import gov.nih.nci.bento.constants.Const;
import gov.nih.nci.bento.model.AbstractPrivateESDataFetcher;
import gov.nih.nci.bento.model.search.MultipleRequests;
import gov.nih.nci.bento.model.search.filter.DefaultFilter;
import gov.nih.nci.bento.model.search.filter.FilterParam;
import gov.nih.nci.bento.model.search.mapper.TypeMapperImpl;
import gov.nih.nci.bento.model.search.mapper.TypeMapperService;
import gov.nih.nci.bento.model.search.query.QueryParam;
import gov.nih.nci.bento.model.search.yaml.YamlQueryFactory;
import gov.nih.nci.bento.service.ESService;
import graphql.schema.idl.RuntimeWiring;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.opensearch.action.search.SearchRequest;
import org.opensearch.client.Request;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.util.*;

import static graphql.schema.idl.TypeRuntimeWiring.newTypeWiring;

@Component
public class PrivateESDataFetcher extends AbstractPrivateESDataFetcher {
    private static final Logger logger = LogManager.getLogger(PrivateESDataFetcher.class);
    private final YamlQueryFactory yamlQueryFactory;
    private final TypeMapperService typeMapper = new TypeMapperImpl();

    public PrivateESDataFetcher(ESService esService) {
        super(esService);
        yamlQueryFactory = new YamlQueryFactory(esService);
    }

    @Override
    public RuntimeWiring buildRuntimeWiring() throws IOException {
        return RuntimeWiring.newRuntimeWiring()
                .type(newTypeWiring("QueryType")
                        .dataFetchers(yamlQueryFactory.createYamlQueries(Const.ES_ACCESS_TYPE.PRIVATE))
                        .dataFetcher("idsLists", env -> idsLists())
                        .dataFetchers(yamlQueryFactory.createYamlQueries(Const.ES_ACCESS_TYPE.PUBLIC))
                        .dataFetcher("numberOfTrials", env -> getNodeCount(TRIALS_COUNT_END_POINT))
                        .dataFetcher("numberOfSubjects", env -> getNodeCount(SUBJECTS_COUNT_END_POINT))
                        .dataFetcher("numberOfFiles", env -> getNodeCount(FILES_COUNT_END_POINT))
                )
                .build();
    }




    private Map<String, Object> idsLists() throws IOException {
        final String[][] PROPERTIES = new String[][]{
                new String[]{"subject_id", "subject_ids"}
        };
        Request request = new Request("GET", SUBJECTS_END_POINT);
        List<Map<String, Object>> response = esService.collectPage(request, new HashMap<String, Object>(), PROPERTIES, ESService.MAX_ES_SIZE, 0);
        ArrayList<String> subject_ids = new ArrayList<>();
        response.forEach(x -> subject_ids.add((String) x.get("subject_id")));
        Map<String, Object> result = Map.of(
                "subject_ids", subject_ids
        );
        return result;
    }
}